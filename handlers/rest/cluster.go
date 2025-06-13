package rest

import (
	"net/http"

	"github.com/aarthikrao/timeMachine/components/consensus"
	"github.com/aarthikrao/timeMachine/components/dht"
	"github.com/aarthikrao/timeMachine/models/config"
	"github.com/aarthikrao/timeMachine/process/nodemanager"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type clusterMessage struct {
	NodeID      string `json:"node_id,omitempty" bson:"node_id,omitempty"`
	RaftAddress string `json:"raft_address,omitempty" bson:"raft_address,omitempty"`
}

type clusterRestHandler struct {
	cp      consensus.Consensus
	appDht  dht.DHT
	nodeMgr *nodemanager.NodeManager
	log     *zap.Logger
}

func CreateClusterRestHandler(
	cp consensus.Consensus,
	appDht dht.DHT,
	nodeMgr *nodemanager.NodeManager,
	log *zap.Logger,
) *clusterRestHandler {
	return &clusterRestHandler{
		cp:      cp,
		appDht:  appDht,
		nodeMgr: nodeMgr,
		log:     log,
	}
}

func (crh *clusterRestHandler) GetStats(c *gin.Context) {
	c.JSON(http.StatusOK, crh.cp.Stats())
}

func (crh *clusterRestHandler) Join(c *gin.Context) {
	var cm clusterMessage
	c.BindJSON(&cm)

	if cm.NodeID == "" || cm.RaftAddress == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid nodeID or address"})
		return
	}

	if !crh.cp.IsLeader() {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":  "not leader",
				"leader": crh.cp.GetLeaderAddress(),
			},
		)
		return
	}

	if err := crh.cp.Join(cm.NodeID, cm.RaftAddress); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (crh *clusterRestHandler) Remove(c *gin.Context) {
	var cm clusterMessage
	c.BindJSON(&cm)

	if cm.NodeID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid nodeID"})
		return
	}

	if !crh.cp.IsLeader() {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":  "not leader",
				"leader": crh.cp.GetLeaderAddress(),
			},
		)
		return
	}

	if err := crh.cp.Remove(cm.NodeID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (crh *clusterRestHandler) Redistribute(c *gin.Context) {

	if !crh.cp.IsLeader() {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":  "not leader",
				"leader": crh.cp.GetLeaderAddress(),
			},
		)
		return
	}

	// TODO: Yet to implement

}

func (crh *clusterRestHandler) Configure(c *gin.Context) {
	if !crh.cp.IsLeader() {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":  "not leader",
				"leader": crh.cp.GetLeaderAddress(),
			},
		)
		return
	}

	var cf config.Config
	if err := c.BindJSON(&cf); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := crh.nodeMgr.InitAppDHT(cf.SlotPerNodeCount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
